package logger

//func ReadFromFile(filename string) string {
//	file, err := ioutil.ReadFile(filename)
//	if err != nil {
//		log.Panic("Failed to log to file", err)
//		panic(err)
//	}
//	return string(file)
//}

////Test for reading
//func ReadFromFile(filename string) string {
//	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
//	if err != nil {
//		log.Panic("Failed to log to file", err)
//		panic(err)
//	}
//	defer func() {
//		err := file.Close()
//		if err != nil {
//			log.Println("Can't close file")
//		}
//	}()
//	data := make([]byte, 64)
//	str := ""
//	for{
//		n, err := file.Read(d\ata)
//		if err == io.EOF{   // если конец файла
//			break           // выходим из цикла
//		}
//		str += string(data[:n])
//	}
//	return str
//}
//
//
//func TestWriteToFile(t *testing.T) {
//	type args struct {
//		filename string
//		data     string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		 {"write to file text test", args{
//			 filename: "test.log",
//			 data:     "test",
//		 }, "test"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			WriteToFile(tt.args.filename, tt.args.data)
//			if got := ReadFromFile(tt.args.filename); got != tt.want{
//				t.Errorf("ReadFromFile() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
