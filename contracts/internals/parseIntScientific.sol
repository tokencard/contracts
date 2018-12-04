pragma solidity ^0.4.25;

/// @title ParseIntScientific provides floating point in scientific notation (e.g. e-5) parsing funcionality.
contract ParseIntScientific {

  byte constant private DASH_ASCII = byte(45); //decimal value of '-'
  byte constant private DOT_ASCII = byte(46); //decimal value of '.'
  byte constant private ZERO_ASCII = byte(48); //decimal value of '0'
  byte constant private NINE_ASCII = byte(57); //decimal value of '9'
  byte constant private E_ASCII = byte(101); //decimal value of 'e'

  function _parseIntScientific(string _a) internal pure returns (uint) {
    return _parseIntScientific(_a, 0);
  }

  function _parseIntScientific(string _a, uint _b) internal pure returns (uint) {

      bytes memory bresult = bytes(_a);
      uint mint = 0;
      uint mintExp = 0;
      bool decimals = false;
      bool exp = false;
      bool minus = false;
      uint expIndex;

      for (uint i=0; i<bresult.length; i++){
          if ((bresult[i] >= ZERO_ASCII) && (bresult[i] <= NINE_ASCII)  && !exp){
              if (decimals){
                 if (_b == 0) break;
                  else _b--;
              }
              mint *= 10;
              mint += uint(bresult[i]) - uint(ZERO_ASCII);
          }
          else if ((bresult[i] >= ZERO_ASCII) && (bresult[i] <= NINE_ASCII) && minus){
              mintExp *= 10;
              mintExp += uint(bresult[i]) - uint(ZERO_ASCII);
          }
          else if (bresult[i] == DOT_ASCII && !decimals){
            decimals = true;
          }
          else if (bresult[i] == DASH_ASCII && !minus && expIndex + 1 == i ) {
              minus = true;
          }
          else if (bresult[i] == E_ASCII && !exp) {
              exp = true;
              expIndex = i;
          }
          else{
              return 666;
          }
      }

      if (exp && !minus)
          return 666;
      else if (exp && minus && i == expIndex + 2)
          return 666;

      if (mintExp > _b)
          mint /= 10**(mintExp - _b);
      else
          mint *= 10**(_b - mintExp);

      return mint;
  }

}
