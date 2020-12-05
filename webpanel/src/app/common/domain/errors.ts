export interface INotification {
  message: string;
  type: NotificationType;
}

export enum NotificationType {
  INFO = 'INFO',
  WARN = 'WARN',
  ERROR = 'ERROR',
  SUCCESS = 'SUCCESS'
}
