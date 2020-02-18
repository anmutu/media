/*
  author='du'
  date='2020/2/18 4:19'
*/
package dpops

func AddVideoDelRec(vid string) error {
	stmtIns, err := dbConn.Prepare("insert into video_del_rec (video_id) values (?)")
	defer stmtIns.Close()
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}
