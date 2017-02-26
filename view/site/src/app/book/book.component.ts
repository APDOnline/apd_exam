import { Component, OnInit } from '@angular/core';
import {BookService} from "./book.service";
import {Book} from "../model/Book";


@Component({
  selector: 'app-book',
  templateUrl: './book.component.html',
  styleUrls: ['./book.component.css']
})
export class BookComponent implements OnInit {
  error:string;
  books:Book[];
  constructor(private bookService: BookService) {
  }

  ngOnInit():void {
    this.bookService.getList()
      .then(list => {
        this.books = list;

        for(let book of this.books) {
          book.created_at_display = (new Date(book.created_at).toLocaleDateString());
          book.updated_at_display = (new Date(book.updated_at).toLocaleDateString());
        }

      })
      .catch(error => {
        console.error(error);
        this.error = error;
      })
  }

  selectBook(book: Book) {
    console.error(book);
  }

}
