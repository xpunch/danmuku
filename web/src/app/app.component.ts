import { Component, OnInit } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { CommentService, Comment } from './comment.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [WebsocketService, CommentService]
})
export class AppComponent implements OnInit {

  private comment: Comment = {
    text: '',
    color: 'white',
    size: 1,
    position: 0,
    time: 0,
  }
  public visible: boolean = true;
  public opacity: number = 100;

  constructor(private readonly commentService: CommentService) {
    commentService.comments.subscribe(msg => {
      console.log("Response from websocket: ", msg);
      $('#danmu').danmu("addDanmu", msg);
    });
  }

  ngOnInit(): void {
    jQuery("#danmu").danmu({
      left: 0,
      top: 0,
      height: "100%",
      width: "100%",
      speed: 20000,
      opacity: 1,
      font_size_small: 24,
      font_size_big: 48,
      top_botton_danmu_time: 6000
    });
    jQuery('#danmu').danmu('danmuResume');
  }

  start() {
    jQuery('danmu').danmu('danmuStart');
  }

  pause() {
    jQuery('#danmu').danmu('danmuPause');
  }

  resume() {
    jQuery('#danmu').danmu('danmuResume');
  }

  stop() {
    jQuery('#danmu').danmu('danmuStop');
  }

  getime() {
    alert(jQuery('#danmu').data("nowTime"));
  }

  getpaused() {
    alert(jQuery('#danmu').data("paused"));
  }

  // //发送弹幕，使用了文档README.md第7节中推荐的方法
  send() {
    if (!this.comment.text) return
    this.comment.time = jQuery('#danmu').data("nowTime") + 1
    this.commentService.comments.next(this.comment);
    this.comment.text = "";
  }

  opacityChanged() {
    jQuery('#danmu').danmu("setOpacity", this.opacity / 100);
  }

  visibleChanged() {
    if (this.visible) {
      jQuery("#danmu").danmu("setOpacity", this.opacity / 100)
    } else {
      jQuery("#danmu").danmu("setOpacity", 0)
    }
  }
}