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
  short_answer: string;
  choice_len: number;
  question_len: number;
  read_load: number;
  created_at: Date;
  created_at_display: string;
  updated_at: Date;
  updated_at_display: string;
  isShow: boolean;
}

