import { Component } from '@angular/core';
import { Note } from '../note';
import { NotesService } from '../notes.service';

@Component({
  selector: 'app-saved',
  templateUrl: './saved.component.html',
  styleUrls: ['./saved.component.css']
})

export class SavedComponent {
  notes: Note[];
  constructor() {
    let service = new NotesService();
    this.notes = service.getNotes();
  }

  getNotes() {
    return this.notes;
  }
  
}
