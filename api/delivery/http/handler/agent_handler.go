package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"encoding/json"
	"github.com/TenaHub/api/entity"
	"fmt"
	"github.com/TenaHub/api/agent"
)

type AgentHandler struct {
	agentService agent.AgentService
}
func NewAgentHandler(adm agent.AgentService) *AgentHandler {
	return &AgentHandler{agentService: adm}
}


func (adm *AgentHandler) GetSingleAgent(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {
	fmt.Println("Before")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	agent, errs := adm.agentService.Agent(uint(id))
	fmt.Println("after", agent)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agent, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (adm *AgentHandler) GetAgents(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {

	agents, errs := adm.agentService.Agents()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agents, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (adm *AgentHandler) DeleteAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.agentService.DeleteAgent(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (adm *AgentHandler) PostAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//data, _ := ioutil.ReadAll(r.Body)
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	agentData := &entity.Agent{}

	err := json.Unmarshal(body, agentData)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.agentService.StoreAgent(agentData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	return
}
func (adm *AgentHandler) PutAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	agentData, errs := adm.agentService.Agent(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &agentData)
	agentData, errs = adm.agentService.UpdateAgent(agentData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agentData, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}




