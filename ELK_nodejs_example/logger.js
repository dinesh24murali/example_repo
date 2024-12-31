const winston = require('winston');
const dotenv = require('dotenv')
const { ecsStringify, ecsFormat } =  require('@elastic/ecs-winston-format');
dotenv.config()

const winstonLogger = winston.createLogger({
  format: ecsFormat(), 
  transports: [
    new winston.transports.Console()
  ]
});

const logger = {
  info: (msg) => {
    winstonLogger.info(msg)
  },
  error: (msg) => {
    winstonLogger.error(msg)
  },
}


module.exports = logger;
