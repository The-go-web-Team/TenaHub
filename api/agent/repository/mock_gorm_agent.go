package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/agent"
	"errors"
)

type MockAgentGormRepo struct {
	conn *gorm.DB
}

func NewMockAgentGormRepo(db *gorm.DB) agent.AgentRepository{
	return &MockAgentGormRepo{conn:db}
}

func (adm *MockAgentGormRepo) Agent(id uint) (*entity.Agent, []error) {
	agent := entity.MockAgent
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &agent, nil
}
func (adm *MockAgentGormRepo) Agents() ([]entity.Agent, []error) {
	var agents []entity.Agent
	agents = append(agents, entity.MockAgent,entity.MockAgent)
	return agents, nil
}
func (adm *MockAgentGormRepo) UpdateAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := agentData
	return agent, nil

}
func (adm *MockAgentGormRepo) StoreAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := agentData
	return agent, nil
}
func (adm *MockAgentGormRepo) DeleteAgent(id uint) (*entity.Agent, []error) {
	agent, errs := adm.Agent(id)
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return agent, errs
}
