import React from "react";

const CustomButton = (props) =>{
   return( 
   <div>
    <button className = 'button' onClick = {props.onClick} > {props.text} </button>

    </div>
   )
}
export default CustomButton;