import { EntityRepository, Repository } from 'typeorm';
import { User } from '../entities/user.entity.js';

@EntityRepository(User)
export class UserRepository extends Repository<User> {}