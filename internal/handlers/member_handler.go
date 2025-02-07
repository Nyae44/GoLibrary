package handlers

import (
	"encoding/json"
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/services"
	"net/http"
	"strconv"
)

type MemberHandler struct {
	memberService services.MemberService
}

func NewMemberHandler(memberService services.MemberService) *MemberHandler {
	return &MemberHandler{memberService: memberService}
}

//HandleCreateMember to handle requests to create a new member

func (h *MemberHandler) HandleCreateMember(w http.ResponseWriter, r *http.Request) {
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdMember, err := h.memberService.CreateMember(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdMember)
	if err != nil {
		return
	}

}

//HandleUpdateMember to handle requests to update member details

func (h *MemberHandler) HandleUpdateMember(w http.ResponseWriter, r *http.Request) {
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedMember, err := h.memberService.UpdateMember(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedMember)
	if err != nil {
		return
	}
}

//HandleGetMemberByID to handle requests to get individual members

func (h *MemberHandler) HandleGetMemberByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
	}
	member, err := h.memberService.GetMemberByID(string(rune(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(member)
	if err != nil {
		return
	}

}

//HandleListMembers to handles requests to get all members

func (h *MemberHandler) HandleListMembers(w http.ResponseWriter, r *http.Request) {
	member, err := h.memberService.ListMembers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(member)
	if err != nil {
		return
	}
}
