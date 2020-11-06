import {strings} from './localizer';

export const formatter = {
  format_time: function(time_str) {
    let before = new Date(time_str);
    let now = new Date();
    let next_day = new Date(new Date().setHours(0, 0, 0, 0)) / 1000 + 3600 * 24;
    let sec = 0;
    if (before.getTime() / 1000 > next_day) {
      let r = formatter.format_date(time_str);
      return r;
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
    let next_day = new Date(new Date().setHours(0, 0, 0, 0)) / 1000 + 3600 * 24;
    let now = new Date();
    if (before > next_day) {
      before = new Date(before.setHours(0, 0, 0, 0));
    }
    let sec = (now - before) / 1000;
    if (sec < 3600 * 24) {
      return strings.today;
    } if (sec >= 3600 * 24 && sec <= 3600 * 24 * 2) {
      return strings.yesterday;
    }
    return strings.formatString(strings.at, {
      years: before.getFullYear(),
      months: before.getMonth() + 1,
      days: before.getDate(),
    });
  },
};
