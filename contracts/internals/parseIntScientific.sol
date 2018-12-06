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
      uint mint = 0; //the final uint returned
      uint mintDec = 0; //the uint following the decimal point
      uint mintExp = 0; //the exponent
      uint decMinted = 0; //how many decimals were 'minted'.
      uint expIndex = 0; //the position in the byte array that 'e' was found (if found)
      uint shifts; //how many times the final number has to be shofted (left or right) i.e. 10^shifts
      bool decimals = false; //indicates a decimal number, set to true if '.' is found
      bool exp = false; //indicates if the number being parsed has an exponential representation
      bool minus = false; //indicated if the exponent is negative
      bool plus = false; //indicated if the exponent is positive

      for (uint i=0; i<inBytes.length; i++){
          if ((inBytes[i] >= ZERO_ASCII) && (inBytes[i] <= NINE_ASCII)  && !exp){
            //'e' not encountered yet, minting integer part or decimals
                if (decimals){
                  //'.' encountered
                  mintDec *= 10;
                  mintDec += uint(inBytes[i]) - uint(ZERO_ASCII);
                  decMinted++;//keep track of how many decimals the input number had
                }
                else{
                  //integer part (before '.')
                  mint *= 10;
                  mint += uint(inBytes[i]) - uint(ZERO_ASCII);
                }
          }
          else if ((inBytes[i] >= ZERO_ASCII) && (inBytes[i] <= NINE_ASCII) && (minus || plus)){
              //exponential notation (e-/+) has been detected, mint the exponent
              mintExp *= 10;
              mintExp += uint(inBytes[i]) - uint(ZERO_ASCII);
          }
          else if (inBytes[i] == DOT_ASCII){
            //an extra decimal point makes the format invalid
            require(!decimals,"duplicate decimal point");
            decimals = true;
          }
          else if (inBytes[i] == DASH_ASCII) {
            //an extra '-' should be considered an invalid character
            require(!minus,"duplicate -");
            require(expIndex + 1 == i,"- sign not immediately after e");
            minus = true;
          }
          else if (inBytes[i] == PLUS_ASCII) {
            //an extra '+' should be considered an invalid character
            require(!plus,"duplicate +");
            require(expIndex + 1 == i,"+ sign not immediately after e");
            plus = true;
          }
          else if (inBytes[i] == E_ASCII) {
            //an extra 'e' should be considered an invalid character
            require(!exp,"duplicate e");
            exp = true;
            expIndex = i;
          }
          else{
              revert("invalid digit");
          }
      }

      if (exp && !minus && !plus)
        //e12 is a valid notation FIX ME FIX ME FIX ME FIX ME FIX ME FIX ME
        revert("exponent not followed by sign");
      else if (exp && minus && i == expIndex + 2)
        //end of string (e-) without specifying the exponent
        revert("exponent not specified");
      else if (exp && plus && i == expIndex + 2)
        //end of string (e+) without specifying the exponent
        revert("exponent not specified");

      if (minus){
        //e^(-x)
        if (mintExp >= magnitudeMult){
          //the (negative) exponent is bigger than the given parameter for "shifting left".
          //use integer division to reduce the precision.
          mint /= 10**(mintExp - magnitudeMult);
        }
        else{
          //the (negative) exponent is smaller than the given parameter for "shifting left".
          shifts = magnitudeMult - mintExp;
          if (shifts >= decMinted){
            //the decimals are fewer or equal than the shifts: use all of them
            ////shift number and add the decimals at the end
            mint *= 10**(decMinted);
            mint += mintDec;
            //add zeros at the end if needed
            mint *= 10**(shifts - decMinted);
          }
          else{
            //the decimals are more than the shifts
            //use only the ones needed, discard the rest
            mintDec /= 10**(decMinted-shifts);
            //shift number and add the decimals at the end
            mint *= 10**(shifts);
            mint += mintDec;
          }
        }
      }
      else{
        //e^(+x), positive exponent (or not)
        //just shift left as many times as indicated by the exponent and the shift parameter
        shifts = magnitudeMult + mintExp;
        //include decimals if present in the original input
        mint *= 10**(decMinted);
        mint += mintDec;
        //'shift' again if the decimals were fewer that the combined shofts
        mint *= 10**(shifts - decMinted);
      }

      return mint;
  }

}
