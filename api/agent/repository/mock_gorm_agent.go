package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/agent"
	"errors"
	"github.com/TenaHub/api/delivery/http/handler"
)

type MockAgentGormRepo struct {
	conn *gorm.DB
}

func NewMockAgentGormRepo(db *gorm.DB) agent.AgentRepository{
	return &MockAgentGormRepo{conn:db}
}

func (adm *MockAgentGormRepo) AgentById(id uint) (*entity.User, []error) {
	agent := entity.MockUser
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &agent, nil
}

func (adm *MockAgentGormRepo) Agent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := entity.Agent{}
	errs := adm.conn.Select("password").Where("email = ? ", agentData.Email).First(&agent).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	same := handler.VerifyPassword(agentData.Password, agent.Password)
	if same {
		errs := adm.conn.Where("email = ?", agentData.Email).First(&agent).GetErrors()
		return &agent, errs
	}
	return nil, errs

}
func (adm *MockAgentGormRepo) Agents() ([]entity.User, []error) {
	var agents []entity.User
	agents = append(agents, entity.MockUser,entity.MockUser)
	return agents, nil
}
func (adm *MockAgentGormRepo) UpdateAgent(agentData *entity.User) (*entity.User, []error) {
	agent := agentData
	return agent, nil

}
func (adm *MockAgentGormRepo) StoreAgent(agentData *entity.User) (*entity.User, []error) {
	agent := agentData
	return agent, nil
}
func (adm *MockAgentGormRepo) DeleteAgent(id uint) (*entity.User, []error) {
	agent, errs := adm.AgentById(id)
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return agent, errs
}
