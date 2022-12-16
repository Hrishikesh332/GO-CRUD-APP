package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type  Movie struct{
	ID string `json:"id"`
	Lastname string `json:"lastname"`

}
type Director{


}