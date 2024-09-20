package db

import (
	"concurrency/db/compute"
	"concurrency/db/storage"
	"fmt"
	"go.uber.org/zap"
)

type Pipeline struct {
	Engine storage.Storage
}

func NewPipeline(engine storage.Storage) *Pipeline {
	return &Pipeline{
		Engine: engine,
	}
}

func (p *Pipeline) Process(input string) (string, error) {
	Logger.Info("Received input", zap.String("input", input))

	cmd, err := compute.Parse(input)
	if err != nil {
		Logger.Error("Parsing error", zap.Error(err))
		return "", err
	}
	Logger.Info("Parsed command", zap.String("action", cmd.Action), zap.Strings("args", cmd.Args))

	switch cmd.Action {
	case "SET":
		p.Engine.Set(cmd.Args[0], cmd.Args[1])
		return "OK", nil

	case "GET":
		value, err := p.Engine.Get(cmd.Args[0])
		if err != nil {
			Logger.Error("Engine GET error", zap.Error(err))
			return "", err
		}
		return value, nil

	case "DEL":
		p.Engine.Del(cmd.Args[0])
		return "OK", nil

	default:
		Logger.Error("Unknown command", zap.String("action", cmd.Action))
		return "", fmt.Errorf("unknown command: %s", cmd.Action)
	}
}
