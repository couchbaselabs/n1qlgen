package main

import "github.com/sirupsen/logrus"

func setLogLevel(logLevel string) error {
   lvl, err := logrus.ParseLevel(logLevel)
   if err != nil {
     return err
   }
   logrus.SetLevel(lvl)
   return nil
}

func scopeEntry(scope string) *logrus.Entry{
  return logrus.WithFields(logrus.Fields{"scope": scope})
}
func logInfo(scope string, query string){
   scopeEntry(scope).Info(query)
}
func logWarn(scope string, query string){
   scopeEntry(scope).Warn(query)
}
func logDebug(scope string, query string){
   scopeEntry(scope).Debug(query)
}
