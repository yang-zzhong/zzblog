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
      style={{
        color: 'var(--blog-a-color)',
        textDecoration: 'none',
      }}
      href={props.href} 
      onClick={onClick} 
      className={props.className}>{props.children}</Link>
  );
}
