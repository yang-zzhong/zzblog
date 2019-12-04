
import LocalizedStrings from 'react-localization';
 
export let strings = new LocalizedStrings({en: {}});

function notify(lang) {
  window.dispatchEvent(new CustomEvent('lang-changed', {detail: lang}));
}

export const localizer = {
  init: function() {
    return fetch('/langs.json').then(res => {
      if (res.status === 200) {
        return res.json().then(langs => {
          console.log(langs);
          strings.setContent(langs || {});
          let lst = [];
          for(let i in langs) {
            lst.push({name: i, label: langs[i].__label || i});
          }
          return new Promise(r => r(lst));
        }).catch(() => {
          return new Promise(r => r([]));
        });
      }
      return new Promise(r => r([]));
    });
  },
  guess: function() {
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
