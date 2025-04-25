import { Router } from "express"
import authRoute from "./auth"

const apiVersionRouting: Router = Router()

apiVersionRouting.use('/auth', authRoute)


export default apiVersionRouting