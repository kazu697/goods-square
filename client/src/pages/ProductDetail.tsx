import React from 'react';
import { IconButton, makeStyles } from '@material-ui/core';
import TwitterIcon from '@material-ui/icons/Twitter';

const useStyles = makeStyles({
  root: {
    display: 'flex',
    padding: '10px',
    margin: '0 20%',
    gap: '30px',
  },
  rightSide: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    width: '20%',
  },
  image: {
    width: '80%',
  },
});

const ProductDetail: React.FC = () => {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <img
        src="https://1.bp.blogspot.com/-tVeC6En4e_E/X96mhDTzJNI/AAAAAAABdBo/jlD_jvZvMuk3qUcNjA_XORrA4w3lhPkdQCNcBGAsYHQ/s1048/onepiece01_luffy.png"
        alt="luffy"
        className={classes.image}
      />
      <div className={classes.rightSide}>
        <div>
          <h2>グッズ</h2>
          <div>クリアチャーム</div>
        </div>
        <div>
          <h2>選手</h2>
          <div>小泉佳穂</div>
        </div>
        <IconButton>
          <TwitterIcon />
        </IconButton>
      </div>
    </div>
  );
};

export default ProductDetail;
