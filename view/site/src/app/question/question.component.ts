import {Component, OnInit, Input, Output, EventEmitter} from '@angular/core';
import {ActivatedRoute, Params} from '@angular/router';
import {Location} from '@angular/common';
import {Question} from "../model/Question";
import {QuestionService} from "./question.service";
import {SelectItem,Message} from 'primeng/primeng'
import 'rxjs/add/operator/switchMap';
import {Book} from "../model/Book";

@Component({
  selector: 'question-list',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.scss']
})
export class QuestionComponent implements OnInit {
  @Input() book: Book;
  @Input() isExamQuestionSelection: boolean;
  @Output() addingQuestion = new EventEmitter<Question>();

  error: string;
  questions: Question[];
  questionType: SelectItem[];
  questionDifficulty: SelectItem[];
  questionSelection: Question;
  displayDetail: boolean;
  message: Message[] = [];

  constructor(private questionService: QuestionService,
              private route: ActivatedRoute,
              private location: Location) {
  }

  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => {
        if (this.book == undefined) {
          this.book.id = +params['bookId'];
          this.book.name = params['bookName'];
        }
        return this.questionService.getList(this.book.id)
      })
      .subscribe(list => {
        this.questions = list;
      });

    this.questionType = [];
    this.questionType.push({label: 'All', value: null});
    this.questionType.push({label: 'T/F', value: 'T/F'});
    this.questionType.push({label: 'MC', value: 'MC'});
    this.questionType.push({label: 'SA', value: 'SA'});

    this.questionDifficulty = [];
    this.questionDifficulty.push({label: 'All', value: null});
    this.questionDifficulty.push({label: 'EASY', value: 'EASY'});
    this.questionDifficulty.push({label: 'MEDIUM', value: 'MEDIUM'});
    this.questionDifficulty.push({label: 'HARD', value: 'HARD'});

  }

  onRowSelect(event) {
    this.displayDetail = true;
  }

  onAddingQuestion() {
    this.addingQuestion.emit(this.questionSelection);
    for(let i=0; i<this.questions.length; i++) {
      if(this.questions[i].id == this.questionSelection.id) {
        this.questions.splice(i, 1);
        break;
      }
    }
    this.displayDetail = false;
    this.message.push({
      severity: 'info',
      summary: "Question Added",
      detail: `Added question ID: ${this.questionSelection.id}`
    })
  }

  goBack(): void {
    this.location.back();
  }


}
