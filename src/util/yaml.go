package util

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
)

func YamlLoad(fn string, out interface{}) error {
    content, err := ioutil.ReadFile(fn)

    if err != nil {
        log.Println("failed reading yaml file", fn)
        return err
    }

    if err := yaml.Unmarshal(content, out); err != nil {
        log.Println("failed to decode yaml file", fn, err)
        return err
    }

    return nil
}

func YamlStore(fn string, in interface{}) error {
    content, err := yaml.Marshal(in)
    if err != nil {
        log.Println("failed to encode yaml file", fn)
        return err
    }

    return ioutil.WriteFile(fn, content, 0644)
}
