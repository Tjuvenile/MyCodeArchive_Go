package dcs

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/logging"
	"MyCodeArchive_Go/utils/strings_"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func CreateRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"Name": "relation2",
	}, CreateRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(RelationParam)

	err = checkNameParam(params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, false)
	if err != nil {
		return err
	}

	relation := DcsRelations{
		UUID:             uuid.NewString(),
		Name:             params.Name,
		MasterPool:       params.MasterPool,
		SlavePool:        params.SlavePool,
		MasterResourceId: params.MasterResource.ID,
		SlaveResourceId:  params.SlaveResource.ID,
		ResourceType:     params.ResourceType,
		StrategyIds:      strings.Join(params.StrategyIds, ","),
		LastSyncTime:     params.LastSyncTime,
		LastSyncSnap:     params.LastSyncSnap,
		State:            params.Status,
		RunningState:     params.RunningState,
		HealthState:      params.HealthState,
		DataState:        params.DataState,
		Role:             params.Role,
		IsConfigSync:     params.IsConfigSync,
	}
	if err = relation.Create(); err != nil {
		return err
	}
	fmt.Println("end craete raltions")
	return nil
}

func UpdateRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID":    "dcba7a69-4991-45d4-953c-47fbf10a7f66",
		"Name":    "",
		"NewName": "relationNewName3",
	}, UpdateRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(RelationUpdateParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, true)
	if err != nil {
		return err
	}

	var relation DcsRelations
	relation.Name = params.Name
	relation.UUID = params.UUID
	update := map[string]interface{}{
		"name": params.NewName,
	}
	if err = relation.Update(update); err != nil {
		return err
	}
	fmt.Println("end update raltions")
	return nil
}

func DeleteRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "",
		"Name": "relationNewName3",
	}, DeleteRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(BaseParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, true)
	if err != nil {
		return err
	}

	var relation DcsRelations
	relation.Name = params.Name
	relation.UUID = params.UUID
	if err = relation.Delete(); err != nil {
		return err
	}
	fmt.Println("end delete raltions")
	return nil
}

func ShowRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "",
		"Name": "relation2",
	}, ShowRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(BaseParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, true)
	if err != nil {
		return err
	}

	relation := DcsRelations{UUID: params.UUID, Name: params.Name}
	if len(params.UUID) != 0 {
		err = relation.QueryById()
	} else {
		err = relation.QueryByName()
	}
	if err != nil {
		return err
	}
	logging.Log.Infof("result: %+v", relation)
	return nil
}

func ListRelationsExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{}, ListRelationsFun)
	if err != nil {
		return err
	}
	params := anyParams.(FilterParam)

	err = CheckSearchParam(params.SearchParam)
	if err != nil {
		return err
	}

	var relationDb DcsRelations
	relations, total, err := relationDb.List(params.FilterBy, params.FilterValue, params.Order, params.SortBy, params.PageSize, params.PageNumber)
	if err != nil {
		return err
	}

	logging.Log.Infof("%+v", relations)
	logging.Log.Infof("lenth: %d", total)
	return nil
}

// TODO 还未验证
func ListRelationsStrategiesExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{}, ListRelationsStrategiesFun)
	if err != nil {
		return err
	}
	params := anyParams.(FilterParam)

	err = CheckSearchParam(params.SearchParam)
	if err != nil {
		return err
	}

	var relationDb DcsRelations
	relations, total, err := relationDb.List(params.FilterBy, params.FilterValue, params.Order, params.SortBy, params.PageSize, params.PageNumber)
	if err != nil {
		return err
	}

	var strategyDb DcsStrategies
	relationsStrategies := []RelationsStrategies{}
	for _, relation := range relations {
		strategies, err := strategyDb.QueryByIds(strings.Split(relation.StrategyIds, ","))
		if err != nil {
			return err
		}
		for index, strategy := range strategies {
			strategies[index].Description = strings_.UnquoteString(strategy.Description)
		}
		relationsStrategies = append(relationsStrategies, RelationsStrategies{
			DcsRelations: relation,
			Strategies:   strategies,
		})
	}

	logging.Log.Infof("%+v", relations)
	logging.Log.Infof("lenth: %d", total)
	return nil
}

func CreateStrategyExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"Name": "strategy1",
	}, CreateStrategyFun)
	if err != nil {
		return err
	}
	params := anyParams.(StrategyParam)

	err = checkNameParam(params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, StrategyModule, false)
	if err != nil {
		return err
	}

	strategy := DcsStrategies{
		UUID:         uuid.NewString(),
		Name:         params.Name,
		TimePoint:    params.TimePoint,
		Interval:     params.Interval,
		StrategyType: params.StrategyType,
		Description:  params.Description,
		CreateTime:   time.Time{},
	}
	if err = strategy.Create(); err != nil {
		return err
	}

	return nil
}

func DeleteStrategyExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "",
		"Name": "relation2",
	}, DeleteStrategyFun)
	if err != nil {
		return err
	}
	params := anyParams.(BaseParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, StrategyModule, true)
	if err != nil {
		return err
	}

	var strategy DcsStrategies
	strategy.Name = params.Name
	strategy.UUID = params.UUID
	if err = strategy.Delete(); err != nil {
		return err
	}
	return nil
}

func UpdateStrategyExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "",
		"Name": "relation2",
	}, UpdateStrategyFun)
	if err != nil {
		return err
	}
	params := anyParams.(StrategyUpdateParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkNameParam(params.NewName)
	if err != nil {
		return err
	}

	var strategy DcsStrategies
	strategy.UUID = params.UUID
	strategy.Name = params.Name
	update := map[string]interface{}{
		"description": params.Description,
		"name":        params.NewName,
	}
	err = strategy.Update(update)
	if err != nil {
		return err
	}
	return nil
}

func ListStrategiesExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{}, ListStrategiesFun)
	if err != nil {
		return err
	}
	params := anyParams.(FilterParam)

	var strategyDb DcsStrategies
	strategies, total, err := strategyDb.List(params.FilterBy, params.FilterValue, params.Order, params.SortBy, params.PageSize, params.PageNumber)
	if err != nil {
		return err
	}

	logging.Log.Infof("%+v", strategies)
	logging.Log.Infof("lenth: %d", total)
	return nil
}
