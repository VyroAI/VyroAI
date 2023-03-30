package userRepository

//func (ur *UserRepository) GetSelf(ctx context.Context, userID int64) (*entites.User, error) {
//	ctx, span := ur.tracer.Start(ctx, "get-self")
//	defer span.End()
//
//	user, err := ur.database.GetUserByEmail(ctx, email)
//	if err != nil {
//		return nil, err
//	}
//
//	return userEmailDbToModel(user), nil
//}
