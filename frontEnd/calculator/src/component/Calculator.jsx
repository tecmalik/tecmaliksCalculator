import React from 'react'
import style from './calculator.module.css'
import CustomButton from '../reusable/CustomButton'



const Calculator = () =>{
    const Value = 0

    return(
        <div id ={ style.wrapper}>
            <div className="screen">

            </div>
            <div className="body">
                <div className="firstRow">
                    <CustomButton text ='1'/>
                    <CustomButton text ='2'/>
                    <CustomButton text ='3'/>
                    <CustomButton text ='+'/>
                </div>
                <div className="secondRow">
                    <CustomButton text ='4'/>
                    <CustomButton text ='5'/>
                    <CustomButton text ='6'/>
                    <CustomButton text ='-'/>
                </div>
                <div className="thirdRow">
                    <CustomButton text ='7'/>
                    <CustomButton text ='8'/>
                    <CustomButton text ='9'/>
                    <CustomButton text ='*'/>
                </div>
                <div className="forthRow">
                <CustomButton text ='7'/>
                    <CustomButton text ='%'/>
                    <CustomButton text ='0'/>
                    <CustomButton text ='/'/>
                </div>
                <div className="lastRow">
                    <CustomButton text ='C'/>
                    <CustomButton text ='='/>
                </div>


            </div>
            



        </div>
    

    )
    

}
export default Calculator