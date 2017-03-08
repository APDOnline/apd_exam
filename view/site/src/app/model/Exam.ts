import {Book} from "./Book";
import {Question} from "./Question";
/**
 * Created by yen-chiehchen on 2/26/17.
 */

export class Exam {
  id: number;
  name: string;
  book: Book;
  questions: Question[];
  question_count: number;
  status: string;
  enable: boolean;
  created_at: Date;
  updated_at: Date;
}
