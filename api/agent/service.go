package agent

import "github.com/TenaHub/api/entity"

type AgentService interface {
	AgentById(id uint) (*entity.User, []error)
	Agents() ([]entity.User, []error)
	Agent(agent *entity.Agent)(*entity.Agent, []error)
	UpdateAgent(user *entity.User) (*entity.User, []error)
	StoreAgent(user *entity.User) (*entity.User, []error)
	DeleteAgent(id uint) (*entity.User, []error)
}
