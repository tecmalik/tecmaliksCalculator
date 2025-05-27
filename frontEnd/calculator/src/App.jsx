
import Calculator from "./component/Calculator"
import router from './routers/routes';
import {RouterProvider} from "react-router"


function App() {

  return (
    <>
    <RouterProvider router ={router}/>
      {/* <Calculator/> */}
    </>
  )
}

export default App
