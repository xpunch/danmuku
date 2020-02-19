import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { map } from 'rxjs/operators';
import { WebsocketService } from './websocket.service';

const SERVER_URL = 'ws://localhost:3000/comment';

export interface Comment {
  color: string
  size: number
  position: number
  text: string
  time: number
}

@Injectable({
  providedIn: 'root'
})
export class CommentService {

  public comments: Subject<Comment>;

  constructor(wsService: WebsocketService) {
    this.comments = <Subject<Comment>>wsService
      .connect(SERVER_URL)
      .pipe(map((response: MessageEvent): Comment => {
        console.log(response);
        let data = JSON.parse(response.data);
        if (data)
          return {
            text: data.text,
            size: data.size,
            color: data.color,
            position: data.position,
            time: data.time
          }
      }));
  }
}
