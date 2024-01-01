package internal

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type SSM struct {
	client ssmiface.SSMAPI
	prefix string
}

func (s *SSM) ListSecretNames(env string) ([]string, error) {
	client, prefix := s.client, fmt.Sprintf("/%s/%s/", s.prefix, env)

	params := &ssm.GetParametersByPathInput{
		Path:           &prefix,
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	result, err := client.GetParametersByPath(params)
	if err != nil {
		fmt.Println("Error retrieving parameters:", err)
		return []string{}, nil
	}

	var names []string
	for _, param := range result.Parameters {
		names = append(names, strings.Replace(*param.Name, prefix, "", -1))
	}

	return names, nil
}

type SecretValue struct {
	Name  string
	Value string
}

func (s *SSM) GetAllSecrets(env string) ([]SecretValue, error) {
	client, prefix := s.client, fmt.Sprintf("/%s/%s/", s.prefix, env)

	params := &ssm.GetParametersByPathInput{
		Path:           &prefix,
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	result, err := client.GetParametersByPath(params)
	if err != nil {
		fmt.Println("Error retrieving parameters:", err)
		return []SecretValue{}, nil
	}

	names := []SecretValue{}
	for _, param := range result.Parameters {
		names = append(names, SecretValue{Name: strings.Replace(*param.Name, prefix, "", -1), Value: *param.Value})
	}

	return names, nil
}

func Sessions() (*session.Session, string) {
	config := GetConfig()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: config.Profile,
		Config: aws.Config{
			Region: &config.Region,
		},
	}))
	return sess, config.Prefix
}

func NewSSMClient() *SSM {
	// Create AWS Session
	sess, prefix := Sessions()

	ssmsvc := &SSM{ssm.New(sess), prefix}
	return ssmsvc
}

type Param struct {
	Name           string
	WithDecryption bool
	ssmsvc         *SSM
}

func (s *SSM) Param(name string) *Param {
	return &Param{
		Name:           name,
		WithDecryption: true,
		ssmsvc:         s,
	}
}

func (p *Param) GetValue() (string, error) {
	ssmsvc := p.ssmsvc.client
	parameter, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &p.Name,
		WithDecryption: &p.WithDecryption,
	})

	if err != nil {
		return "", err
	}

	value := *parameter.Parameter.Value
	return value, nil
}

func (p *Param) PutValue(secret string, overwrite bool) error {
	ssmsvc := p.ssmsvc.client

	description := "Value created by s cli for managing secrets"
	parameterType := ssm.ParameterTypeSecureString

	_, err := ssmsvc.PutParameter(&ssm.PutParameterInput{
		Name:        &p.Name,
		Description: &description,
		Value:       &secret,
		Type:        &parameterType,
		Overwrite:   &overwrite,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *Param) DeleteValue() error {
	ssmsvc := p.ssmsvc.client
	_, err := ssmsvc.DeleteParameter(&ssm.DeleteParameterInput{
		Name: &p.Name,
	})

	if err != nil {
		return err
	}

	return nil
}
