import {strings} from './localizer';

export const formatter = {
  format_time: function(time_str) {
    let before = new Date(time_str);
    let now = new Date();
    let db = new Date(new Date().setHours(0, 0, 0, 0));
    let sec = 0;
    if (before.getTime() < db.getTime()) {
      return formatter.format_date(time_str);
    }
    sec = (now.getTime() - before) / 1000;
    if (sec <= 3600) {
      let r = strings.formatString(strings.minsAgo, {
        mins: Math.floor(sec / 60)
      });
      return r;
    } else if (sec >= 3600 && sec <= 3600 * 24) {
      let r = strings.formatString(strings.hoursAgo, {
        hours: Math.floor(sec / 3600)
      });
      return r;
    }
  },
  format_date: function(time_str) {
    let before = new Date(time_str);
    let bs = before.getTime();
    let tmp = new Date(new Date().setHours(0, 0, 0, 0));
    let before_day =  tmp / 1000 - 3600 * 24;
    let now = new Date();
    if (bs > tmp.getTime()) {
      return strings.today;
    } else if (bs / 1000 > before_day) {
      return strings.yesterday;
    }
    return strings.formatString(strings.at, {
      years: before.getFullYear(),
      months: before.getMonth() + 1,
      days: before.getDate(),
    });
  },
};
