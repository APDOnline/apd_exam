import {Difficulty} from "./Difficulty";
import {Option} from "./Option";
/**
 * Created by yen-chiehchen on 2/23/17.
 */

export class Question {
  id: number;
  type: string;
  text: string;
  image: string;
  category: string;
  difficulty: Difficulty[];
  learning_outcome: string;
  lo_number: string;
  options: Option[];
  reference: string;
  answer_exp: string;
  created_at: string;
  updated_at: string;
}

