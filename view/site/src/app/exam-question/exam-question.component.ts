import { Component, OnInit, Input } from '@angular/core';
import {Question} from "../model/Question";
import {SelectItem,Message} from 'primeng/primeng'
import {Exam} from "../model/Exam";

@Component({
  selector: 'app-exam-question',
  templateUrl: './exam-question.component.html',
  styleUrls: ['./exam-question.component.scss']
})
export class ExamQuestionComponent implements OnInit {
  @Input() exam: Exam;
  questionType: SelectItem[];
  questionDifficulty: SelectItem[];
  questionSelection: Question;
  displayDetail: boolean;

  constructor() { }

  ngOnInit() {
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
}
