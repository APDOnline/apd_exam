import {Difficulty} from "./Difficulty";
import {Option} from "./Option";
import {Reference} from "./Reference";
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
  reference: Reference;
  answer_exp: string;
  short_answer: string;
  choice_len: number;
  question_len: number;
  read_load: number;
  created_at: Date;
  updated_at: Date;
  isShow: boolean;
}

