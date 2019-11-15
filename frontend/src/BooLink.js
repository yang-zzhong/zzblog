import 'boo-route';
import React from 'react';
import Link from '@material-ui/core/Link';

export default function BooLink(props) {
  let onClick = e => {
    e.preventDefault();
    window.boo.location.go(props.href);
  }
  return (
    <Link 
      href={props.href} 
      onClick={onClick} 
      className={props.className}>{props.children}</Link>
  );
}
