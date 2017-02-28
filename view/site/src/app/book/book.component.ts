import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router'
import {BookService} from "./book.service";
import {Book} from "../model/Book";
import {Exam} from "../model/Exam";


@Component({
  selector: 'app-book',
  templateUrl: './book.component.html',
  styleUrls: ['./book.component.css']
})
export class BookComponent implements OnInit {
  error:string;
  books:Book[];
  private showExamCreateForm: boolean;
  private createExamBookTitle: string;
  private inputExamName: string;
  private createExamFromBook: Book;

  constructor(private bookService: BookService, private router: Router) {
  }

  ngOnInit():void {
    this.bookService.getList()
      .then(list => {
        this.books = list;
      })
      .catch(error => {
        console.error(error);
        this.error = error;
      })


  }

  createExamClicked(book: Book) {
    this.createExamBookTitle = book.name;
    this.showExamCreateForm = true;
    this.createExamFromBook = book;
  }

  updateExamName(examName: string) {
    this.inputExamName = examName;
  }

  submitExamCreation() {
    var exam: Exam = new Exam();
    exam.book = this.createExamFromBook;
    exam.name = this.inputExamName;
    this.bookService.submitExamCreation(exam).then(exam => {
      this.router.navigateByUrl(`/exam/${exam.id}`);
    });
  }

}
