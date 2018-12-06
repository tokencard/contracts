pragma solidity ^0.4.25;

/// @title ParseIntScientific provides floating point in scientific notation (e.g. e-5) parsing funcionality.
contract ParseIntScientific {

  byte constant private PLUS_ASCII = byte(43); //decimal value of '+'
  byte constant private DASH_ASCII = byte(45); //decimal value of '-'
  byte constant private DOT_ASCII = byte(46); //decimal value of '.'
  byte constant private ZERO_ASCII = byte(48); //decimal value of '0'
  byte constant private NINE_ASCII = byte(57); //decimal value of '9'
  byte constant private E_ASCII = byte(101); //decimal value of 'e'

  function _parseIntScientific(string inString) internal pure returns (uint) {
    return _parseIntScientific(inString, 0);
  }

  function _parseIntScientific(string inString, uint magnitudeMult) internal pure returns (uint) {

      bytes memory inBytes = bytes(inString);
      uint mint = 0;
      uint mintDec = 0;
      uint mintExp = 0;
      uint decMinted = 0;
      uint expIndex = 0;
      uint shifts;
      bool decimals = false;
      bool exp = false;
      bool minus = false;
      bool plus = false;

      for (uint i=0; i<inBytes.length; i++){
          if ((inBytes[i] >= ZERO_ASCII) && (inBytes[i] <= NINE_ASCII)  && !exp){
                if (decimals){
                  mintDec *= 10;
                  mintDec += uint(inBytes[i]) - uint(ZERO_ASCII);
                  decMinted++;
                }
                else{
                  mint *= 10;
                  mint += uint(inBytes[i]) - uint(ZERO_ASCII);
                }
          }
          else if ((inBytes[i] >= ZERO_ASCII) && (inBytes[i] <= NINE_ASCII) && (minus || plus)){
              mintExp *= 10;
              mintExp += uint(inBytes[i]) - uint(ZERO_ASCII);
          }
          else if (inBytes[i] == DOT_ASCII && !decimals){
            decimals = true;
          }
          else if (inBytes[i] == DASH_ASCII && !minus && expIndex + 1 == i ) {
              minus = true;
          }
          else if (inBytes[i] == PLUS_ASCII && !plus && expIndex + 1 == i ) {
              plus = true;
          }
          else if (inBytes[i] == E_ASCII && !exp) {
              exp = true;
              expIndex = i;
          }
          else{
              revert("wrong format");
          }
      }

      if (exp && !minus && !plus)
        revert("exponent not followed by sign");
      else if (exp && minus && i == expIndex + 2)
        revert("exponent not specified");
      else if (exp && plus && i == expIndex + 2)
        revert("exponent not specified");

      if (minus){
        //e^-x
        if (mintExp >= magnitudeMult){
          mint /= 10**(mintExp - magnitudeMult);
        }
        else{
          shifts = magnitudeMult - mintExp;
          if (shifts >= decMinted){
            mint *= 10**(decMinted);
            mint += mintDec;
            mint *= 10**(shifts - decMinted);
          }
          else{
            mintDec /= 10**(decMinted-shifts);
            mint *= 10**(shifts);
            mint += mintDec;
          }
        }
      }
      else{
        //e^+x
        shifts = magnitudeMult + mintExp;
        mint *= 10**(decMinted);
        mint += mintDec;
        mint *= 10**(shifts - decMinted);
      }

      return mint;
  }

}
