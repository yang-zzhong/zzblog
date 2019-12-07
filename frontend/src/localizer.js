
import LocalizedStrings from 'react-localization';
 
export let strings = new LocalizedStrings({en: {}});

function notify(lang) {
  window.dispatchEvent(new CustomEvent('lang-changed', {detail: lang}));
}

export const localizer = {
  langs: [],
  __ready: false,
  ready: function() {
    if (localizer.__ready) {
      return new Promise(r => r(localizer.langs));
    }
    localizer.__ready = true;
    return localizer.init();
  },
  init: function() {
    return fetch('/langs.json').then(res => {
      if (res.status === 200) {
        return res.json().then(langs => {
          strings.setContent(langs || {});
          let lst = [];
          for(let i in langs) {
            lst.push({name: i, label: langs[i].__label || i});
          }
          localizer.langs = lst;
          return new Promise(r => r(localizer.langs));
        }).catch(() => {
          return new Promise(r => r([]));
        });
      }
      return new Promise(r => r([]));
    });
  },
  guess: function() {
    const ctx = window.boo.location.context;
    if (ctx.query_params.lang) {
      return ctx.query_params.lang;
    }
    const l = localStorage.getItem('lang');
    if (l) {
      return l;
    }
    return localizer.browser();
  },
  use: function(lang) {
    strings.setLanguage(lang);
    localStorage.setItem('lang', lang);
    notify(lang);
  },
  browser: function() {
    if (navigator.appName === 'Netscape') {
      return navigator.language;
    }
    return navigator.browserLanguage;
  },
  lang: function() {
    return strings.getLanguage() || 'en';
  }
};
