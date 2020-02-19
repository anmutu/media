/*
  author='du'
  date='2020/2/18 4:19'
*/
package s_dbops

//读取视频删除纪录表
func ReadVideoDelRec(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_id from video_del_rec limit ?")
	defer stmtOut.Close()

	var ids []string
	if err != nil {
		return ids, nil
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		return ids, nil
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, nil
		}
		ids = append(ids, id)
	}
	return ids, nil
}

//删除视频纪录表
func DeleteVideoDelRec(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_del_rec where video_id=?")
	defer stmtDel.Close()
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}
