package mongos

import (
	"context"
	"time"
	
)

func MongoUsers() repositories.UserRepository {
	return &mongoUsers{}
}

func (m *mongoUsers) ChangeActiveUser(ctx context.Context, inDom *repositories.UserDomain) error {
	query := queries.NewChangeActiveUserQuery()
	params := params.NewChangeActiveUserParams(inDom)
	validator := validators.NewChangeActiveUserValidator()
	result := results.NewChangeActiveUserResult()
	transformer := transformers.NewChangeActiveUserTransformer()

	// Add your code here

	return nil
}

	err := query.Execute(ctx, params, validator, result, transformer)
	if err != nil {
		return err
	}

	return nil
}


func (m *mongoUsers) GetByEmail(ctx context.Context, inDom *repositories.UserDomain) (*repositories.UserDomain, error) {
	query := queries.NewGetByEmailQuery()
	params := params.NewGetByEmailParams(inDom)
	validator := validators.NewGetByEmailValidator()
	result := results.NewGetByEmailResult()
	transformer := transformers.NewGetByEmailTransformer()

	err := query.Execute(ctx, params, validator, result, transformer)
	if err != nil {
		return nil, err
	}

	return result.Domain, nil
}

func (m *mongoUsers) Store(ctx context.Context, inDom *repositories.UserDomain) error {
	query := queries.NewStoreQuery()
	params := params.NewStoreParams(inDom)
	validator := validators.NewStoreValidator()
	result := results.NewStoreResult()
	transformer := transformers.NewStoreTransformer()

	err := query.Execute(ctx, params, validator, result, transformer)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoUsers) Update(ctx context.Context, inDom *repositories.UserDomain) error {
	query := queries.NewUpdateQuery()
	params := params.NewUpdateParams(inDom)
	validator := validators.NewUpdateValidator()
	result := results.NewUpdateResult()
	transformer := transformers.NewUpdateTransformer()

	err := query.Execute(ctx, params, validator, result, transformer)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoUsers) Delete(ctx context.Context, inDom *repositories.UserDomain) error {
	query := queries.NewDeleteQuery()
	params := params.NewDeleteParams(inDom)
	validator := validators.NewDeleteValidator()
	result := results.NewDeleteResult()
	transformer := transformers.NewDeleteTransformer()

	err := query.Execute(ctx, params, validator, result, transformer)
	if err != nil {
		return err
	}

	return nil
}
