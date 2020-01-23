package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/agent"
)

type AgentGormRepo struct {
	conn *gorm.DB
}

func NewAgentGormRepo(db *gorm.DB) agent.AgentRepository{
	return &AgentGormRepo{conn:db}
}

func (adm *AgentGormRepo) Agent(id uint) (*entity.Agent, []error) {
	agent := entity.Agent{}
	errs := adm.conn.First(&agent, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &agent, errs
}
func (adm *AgentGormRepo) Agents() ([]entity.Agent, []error) {
	var agents []entity.Agent
	errs := adm.conn.Find(&agents).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agents, errs
}
func (adm *AgentGormRepo) UpdateAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := agentData
	errs := adm.conn.Save(agent).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs

}
func (adm *AgentGormRepo) StoreAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := agentData
	errs := adm.conn.Create(agent).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
func (adm *AgentGormRepo) DeleteAgent(id uint) (*entity.Agent, []error) {
	agent, errs := adm.Agent(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = adm.conn.Delete(agent, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
