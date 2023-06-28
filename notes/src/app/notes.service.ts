import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Note } from './note';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class NotesService {

  constructor(private httpClient: HttpClient) {

  }

  getNotes(): Observable<Array<Note>> {
    return this.httpClient.get<Array<Note>>('/api/v1/list_notes');
  }

  addNote(note: Note): Observable<Note> {
    return this.httpClient.post<Note>('/api/v1/add_notes', note);
  }

}
