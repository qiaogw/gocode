import { date } from 'quasar'
import cronParser from 'cron-parser'

export const FormatDateTime = (datetime) => {
  return date.formatDate(datetime, 'YYYY-MM-DD HH:mm:ss')
}
export const FormatDate = (datetime) => {
  return date.formatDate(datetime, 'YYYY-MM-DD')
}

export const CronToStruct = (cronString) => {
  try {
    const interval = cronParser.parseExpression(cronString)
    // return {
    //   Seconds: interval.fields[0].range,
    //   Minutes: interval.fields[1].range,
    //   Hours: interval.fields[2].range,
    //   Month: interval.fields[3].range,
    //   DayOfWeek: interval.fields[5].range,
    //   Year: interval.fields[6].range,
    //   DayOfMonth: interval.fields[4].range,
    // }
    const fields = JSON.parse(JSON.stringify(interval.fields))
    return fields
  } catch (error) {
    console.error('Error parsing cron expression:', error)
    return null
  }
}
