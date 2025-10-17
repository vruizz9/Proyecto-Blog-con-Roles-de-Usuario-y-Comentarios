import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Comment {
  id: string;
  blogId: string;
  userId: string;
  content: string;
}

@Injectable({
  providedIn: 'root'
})
export class CommentService {
  private apiUrl = 'http://localhost:3000/comments';

  constructor(private http: HttpClient) {}

  getCommentsByBlogId(blogId: string): Observable<Comment[]> {
    return this.http.get<Comment[]>(`${this.apiUrl}?blogId=${blogId}`);
  }

  addComment(comment: Comment): Observable<Comment> {
  return this.http.post<Comment>(this.apiUrl, comment);
}
}