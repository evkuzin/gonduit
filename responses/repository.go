package responses

import "github.com/evkuzin/gonduit/entities"

// RepositoryQueryResponse is the result of repository.query operations.
type RepositoryQueryResponse []*entities.Repository
