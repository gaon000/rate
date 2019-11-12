import xml.etree.ElementTree as elemtree
tree = elemtree.parse("nine.xml")
root = tree.getroot()
for i in root.iter('Row'):
    print("="*60)
    print(f"모집단위 = {i.findtext('모집단위')}")
    print(f"접수인원 = {i.findtext('접수인원')}")
    print(f"선발예정인원 = {i.findtext('선발예정인원')}")
    print(f"경쟁률 = {i.findtext('경쟁률')}")