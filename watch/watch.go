package watch

import (
	"fmt"
	"github.com/mewa/wuff/config"
	"github.com/mewa/wuff/mail"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Serve(service *config.Service, conf *config.Config) {
	args := strings.Split(service.Check, " ")
	name := args[0]

	retries := 0

	for {
		cmd := exec.Command(name, args[1:]...)
		err := cmd.Start()

		if err != nil {
			log.WithField("service", service.Name).Errorln(err)
		}

		err = cmd.Wait()

		if err != nil {
			log.WithField("service", service.Name).Errorln(err)
			log.WithField("service", service.Name).Errorf("Service %s is down.", service.Name)

			if retries == 0 {
				mail.SendEmail(fmt.Sprintf("Service %s is down.", service.Name), conf)
			} else if retries >= service.Retries {
				mail.SendEmail(fmt.Sprintf("Service %s can't be started after %d attempts.", service.Name, retries), conf)
				break
			}

			retries = retries + 1

			startService(service)
			waitOrYield(service.RetryPeriod)
		} else {
			if retries != 0 {
				mail.SendEmail(fmt.Sprintf("Service %s has been started after %d attempts.", service.Name, retries), conf)
				retries = 0
			}
			waitOrYield(service.CheckPeriod)
		}
	}
}

func startService(service *config.Service) {
	args := strings.Split(service.Start, " ")
	name := args[0]

	cmd := exec.Command(name, args[1:]...)
	err := cmd.Start()

	if err != nil {
		log.WithField("service", service.Name).Errorln(err)
	}

	err = cmd.Wait()

	if err != nil {
		log.WithField("service", service.Name).Errorln(err)
	}
}

func waitOrYield(seconds int) {
	if seconds > 0 {
		time.Sleep(time.Second * time.Duration(seconds))
	} else {
		runtime.Gosched()
	}
}
