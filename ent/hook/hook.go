// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/ent"
)

// The CompanyDetailFunc type is an adapter to allow the use of ordinary
// function as CompanyDetail mutator.
type CompanyDetailFunc func(context.Context, *ent.CompanyDetailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyDetailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompanyDetailMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyDetailMutation", m)
}

// The CompanyEngineerFunc type is an adapter to allow the use of ordinary
// function as CompanyEngineer mutator.
type CompanyEngineerFunc func(context.Context, *ent.CompanyEngineerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyEngineerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompanyEngineerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyEngineerMutation", m)
}

// The CompanyTokenFunc type is an adapter to allow the use of ordinary
// function as CompanyToken mutator.
type CompanyTokenFunc func(context.Context, *ent.CompanyTokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyTokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompanyTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyTokenMutation", m)
}

// The CompanyUserFunc type is an adapter to allow the use of ordinary
// function as CompanyUser mutator.
type CompanyUserFunc func(context.Context, *ent.CompanyUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompanyUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyUserMutation", m)
}

// The JobAuthorFunc type is an adapter to allow the use of ordinary
// function as JobAuthor mutator.
type JobAuthorFunc func(context.Context, *ent.JobAuthorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobAuthorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobAuthorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobAuthorMutation", m)
}

// The JobContractorFunc type is an adapter to allow the use of ordinary
// function as JobContractor mutator.
type JobContractorFunc func(context.Context, *ent.JobContractorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobContractorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobContractorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobContractorMutation", m)
}

// The JobDetailFunc type is an adapter to allow the use of ordinary
// function as JobDetail mutator.
type JobDetailFunc func(context.Context, *ent.JobDetailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobDetailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobDetailMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobDetailMutation", m)
}

// The JobLayerFunc type is an adapter to allow the use of ordinary
// function as JobLayer mutator.
type JobLayerFunc func(context.Context, *ent.JobLayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobLayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobLayerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobLayerMutation", m)
}

// The JobOwnerFunc type is an adapter to allow the use of ordinary
// function as JobOwner mutator.
type JobOwnerFunc func(context.Context, *ent.JobOwnerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobOwnerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobOwnerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobOwnerMutation", m)
}

// The JobPaymentsFunc type is an adapter to allow the use of ordinary
// function as JobPayments mutator.
type JobPaymentsFunc func(context.Context, *ent.JobPaymentsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobPaymentsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobPaymentsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobPaymentsMutation", m)
}

// The JobProgressFunc type is an adapter to allow the use of ordinary
// function as JobProgress mutator.
type JobProgressFunc func(context.Context, *ent.JobProgressMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobProgressFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobProgressMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobProgressMutation", m)
}

// The JobRelationsFunc type is an adapter to allow the use of ordinary
// function as JobRelations mutator.
type JobRelationsFunc func(context.Context, *ent.JobRelationsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobRelationsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobRelationsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobRelationsMutation", m)
}

// The JobSupervisorFunc type is an adapter to allow the use of ordinary
// function as JobSupervisor mutator.
type JobSupervisorFunc func(context.Context, *ent.JobSupervisorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobSupervisorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.JobSupervisorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobSupervisorMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
