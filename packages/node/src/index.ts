import { readFileSync } from 'fs';
import { join } from 'path';
import { GetParametersByPathCommand, SSMClient } from '@aws-sdk/client-ssm';
import { fromIni } from '@aws-sdk/credential-provider-ini';

class Config {
    public prefix = '';
    public profile = '';
    public region = '';

    constructor(config: string) {
        const lines = config.split('\n');
        for (let index = 0; index < lines.length; index++) {
            const element = lines[index];
            const configuration = element!.split(': ');
            const prefix = configuration[0];
            if (
                prefix === 'prefix' ||
                prefix === 'profile' ||
                prefix === 'region'
            ) {
                this[prefix] = configuration[1]!;
            }
        }

        if (this.prefix === '' || this.prefix === '' || this.region === '')
            throw new Error('Unable to parse config');
    }
}

export function s(callback: () => void) {
    const runningDirectory = process.cwd();
    const configPath = join(runningDirectory, '.s.yml');
    const config = new Config(readFileSync(configPath, 'utf-8'));

    const ssmClient = new SSMClient({
        region: config.region,
        credentials: fromIni({ profile: config.profile }),
    });

    ssmClient
        .send(
            new GetParametersByPathCommand({
                Path: `/${config.prefix}/`,
                WithDecryption: true,
            })
        )
        .then((data) => {
            if (data && data.Parameters) {
                for (let index = 0; index < data.Parameters.length; index++) {
                    const param = data.Parameters[index]!;
                    process.env[param.Name!.replace(`/${config.prefix}/`, '')] =
                        param.Value!;
                }
                console.log('s: Environment Variables set!');
            }
            callback();
        });
}
