package seq

import "github.com/google/uuid"

/**
  生成序列
 */
func Generator() string{
	seq := uuid.New()
	return seq.String()
}
