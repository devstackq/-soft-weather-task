package managers

import (
	"context"
	"errors"
	"fmt"
	"soft-weater/internal/clients/http/service-1"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

type Solver interface {
	Do(ctx context.Context, req models.Solve) (int, error)
}

type SolveManager struct {
	accountSrv  service.Service
	taskRepo    drivers.TaskRepository
	historyRepo drivers.HistoryRepository
}

func NewSolver(hRepo drivers.HistoryRepository, tRepo drivers.TaskRepository, accountSrv service.Service) Solver {
	return &SolveManager{
		historyRepo: hRepo,
		taskRepo:    tRepo,
		accountSrv:  accountSrv,
	}
}

const debtLimit = 2000

func (t *SolveManager) Do(ctx context.Context, req models.Solve) (int, error) {

	task, err := t.taskRepo.GetByID(ctx, req.TaskID)
	if err != nil {
		return 0, err
	}

	account, err := t.accountSrv.GetAccount(ctx, req.UserID)
	if err != nil {
		return 0, err
	}
	fmt.Printf("account %+v task %+v req %+v \n", account, task, req)

	if account.Debt > debtLimit || account.Debt+task.Price >= debtLimit { //sql when
		return 0, errors.New("debt more than 1000")
	}

	result := solveTask(req.Input, req.TaskID)

	account.Debt = task.Price
	account.UserID = req.UserID

	if err = t.accountSrv.IncreaseDebt(ctx, account); err != nil {
		return 0, err
	}

	historyReq := models.History{
		UserID: req.UserID,
		TaskID: req.TaskID,
	}

	if err = t.historyRepo.Create(ctx, historyReq); err != nil {
		return 0, err
	}

	return result, nil
}

func solveTask(req any, taskID string) int {

	height := req.([]int)

	switch taskID {
	case "1":
		break
	case "2":
	default:
		fmt.Println("Task ID invalid ", taskID)
	}

	n := len(height)

	if n <= 2 {
		return 0
	}

	left := 0
	right := n - 1
	leftMax := 0
	rightMax := 0
	result := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}
