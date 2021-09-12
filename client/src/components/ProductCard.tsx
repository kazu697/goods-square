import React from 'react';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core';

const useStyles = makeStyles({
  root: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  image: {
    width: '100%',
  },
});

const ProductCard: React.FC = () => {
  const classes = useStyles();
  return (
    <Paper elevation={3} className={classes.root}>
      <img
        src="https://1.bp.blogspot.com/-tVeC6En4e_E/X96mhDTzJNI/AAAAAAABdBo/jlD_jvZvMuk3qUcNjA_XORrA4w3lhPkdQCNcBGAsYHQ/s1048/onepiece01_luffy.png"
        alt="luffy"
        className={classes.image}
      />
      <div>
        <h3>商品名</h3>
        <div>詳細</div>
      </div>
    </Paper>
  );
};

export default ProductCard;
