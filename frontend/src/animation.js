
import {helper} from './helper';

export const animation = {
  fade_in: function(nodes) {
    return {
      keyframes: [
        {opacity: 0},
        {opacity: 1}
      ],
      time: 300,
      nodes: nodes
    };
  },
  fade_out: function(nodes) {
    return {
      keyframes: [
        {opacity: 1},
        {opacity: 0}
      ],
      time: 300,
      nodes: nodes
    };
  },
  scale_in: function(nodes) {
    return {
      keyframes: [
        {transform: 'scale(1.2)', opacity: 0},
        {transform: 'scale(1)', opacity: 1}
      ],
      time: 300,
      nodes: nodes
    };
  },
  scale_out: function(nodes) {
    return {
      keyframes: [
        {transform: 'scale(1)', opacity: 1},
        {transform: 'scale(0.8)', opacity: 0}
      ],
      time: 300,
      nodes: nodes
    };
  },
  left_in: function(nodes) {
    return {
      keyframes: [
        {transform: 'translate(-10px)', opacity: 0},
        {transform: 'translate(0px)', opacity: 1}
      ],
      time: 300,
      nodes: nodes
    };
  },
  right_in: function(nodes) {
    return {
      keyframes: [
        {transform: 'translate(10px)', opacity: 0},
        {transform: 'translate(0px)', opacity: 1}
      ],
      time: 300,
      nodes: nodes
    };
  },
  top_in: function(nodes) {
    return {
      keyframes: [
        {transform: 'translateY(-20px)', opacity: 0},
        {transform: 'translateY(0px)', opacity: 1}
      ],
      time: 300,
      nodes: nodes
    };
  },
  bottom_in: function(nodes) {
    return {
      keyframes: [
        {transform: 'translateY(20px) scale(0.8)', opacity: 0},
        {transform: 'translateY(0px) scale(1)', opacity: 1}
      ],
      time: 300,
      nodes: nodes
    };
  },
  bottom_out: function(nodes) {
    return {
      keyframes: [
        {transform: 'translateY(0px)', opacity: 1},
        {transform: 'translateY(20px)', opacity: 0}
      ],
      time: 300,
      nodes: nodes
    };
  },
  top_out: function(nodes) {
    return {
      keyframes: [
        {transform: 'translateY(0px)', opacity: 1},
        {transform: 'translateY(-20px)', opacity: 0}
      ],
      time: 300,
      nodes: nodes
    };
  },
  right_out: function(nodes) {
    return {
      keyframes: [
        {transform: 'translate(0px)', opacity: 1},
        {transform: 'translate(10px)', opacity: 0}
      ],
      time: 300,
      nodes: nodes
    };
  },
  left_out: function(nodes) {
    return {
      keyframes: [
        {transform: 'translate(0px)', opacity: 1},
        {transform: 'translate(-10px)', opacity: 0}
      ],
      time: 300,
      nodes: nodes
    };
  },
  play: function(conf) {
    if (conf.node && conf.node.animate) {
      if (helper.is_visible(conf.node)) {
        conf.node.animate(conf.keyframes, conf.time);
      }
    } else if (conf.nodes) {
      let play = (nodes, frames, delay, duration) => {
        if (nodes.length === 0) {
          return new Promise(r => r());
        }
        return new Promise(r => {
          let node = nodes[0];
          if (!helper.is_visible(node)) {
            node.style.visibility = 'visible';
            play(nodes.slice(1), frames, delay, duration).then(() => r());
            return;
          }
          setTimeout(() => {
            node.style.visibility = 'visible';
            node.animate(frames, duration);
            play(nodes.slice(1), frames, delay, duration).then(() => r());
          }, delay);
        });
      }
      if (conf.delay && conf.delay > 0) {
        let nodes = [];
        for (let i = 0; i < conf.nodes.length; ++i) {
          conf.nodes[i].style.visibility = 'hidden';
          nodes.push(conf.nodes[i]);
        }
        return play(nodes, conf.keyframes, conf.delay, conf.time);
      }
      conf.nodes.forEach(node => {
        if(node && node.animate && helper.is_visible(node)) {
          node.animate(conf.keyframes, conf.time);
        }
      });
    }
    return new Promise(r =>  {
      setTimeout(() => { r() })
    }, conf.time);
  }
};
