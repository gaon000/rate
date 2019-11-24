import xml.etree.ElementTree as elemtree
import pymongo
conn = pymongo.MongoClient("localhost",27017)
db = conn.officer
collection_nine = db.nine
collection_seven = db.seven
collection_five = db.five

tree = elemtree.parse("seven.xml")
root = tree.getroot()

for i in root.iter('Row'):
	collection_seven.insert({"occupation":i.findtext('모집단위'),"applicant":i.findtext('접수인원'),"selected":i.findtext('선발예정인원'),"rate":i.findtext('경쟁률')})

