import {Book} from "./Book";
import {Question} from "./Question";
/**
 * Created by yen-chiehchen on 2/26/17.
 */

export class Exam {
  id: number;
  name: string;
  book: Book;
  question: Question[];
  status: string;
  enable: boolean;
  created_at: Date;
  created_at_display: string;
  updated_at: Date;
  updated_at_display: string;
}
