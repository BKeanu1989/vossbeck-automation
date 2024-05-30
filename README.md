# Vossbeck automation
Wenn Karten (Deal) in Trello in die Status "Won", abgelegt werden, soll 

---


Wenn Deal Status → “won”, dann

erstelle einen Google Drive Ordner (Name = Dealname) im Shared ordner mit Arbio und fülle alle PDF Anhänge des Deals in den Ordner
erstelle in Trello Arbeitsbereich Vossbeck Management, Board Daily, Liste “Timo Säubert” eine neue Karte (Kartenname = Dealname) und hänge die PDF Dateien aus dem Deal der Karte an


zum testen, einen anderen shared ordner benutzen.


---
## Trello API
https://developer.atlassian.com/cloud/trello/guides/power-ups/building-a-power-up-part-one/


### Glitch
https://glitch.com/pricing



# Trello API
https://developer.atlassian.com/cloud/trello/rest/api-group-actions/#api-group-actions

https://developer.atlassian.com/cloud/trello/guides/power-ups/building-a-power-up-part-one/#:~:text=The%20most%20important%20field%20is,into%20the%20iframe%20connector%20field


---
to search:

google drive playground

--
https://protocoderspoint.com/nodejs-script-to-upload-file-to-google-drive-using-googleapis/


---
google drive api

Das Erstellen von Dienstkontoschlüsseln ist deaktiviert
The organization policy constraint 'iam.disableServiceAccountKeyCreation' is enforced on your organization.
Possible Causes: Your Organization Administrator enforced the Organization Policy to prevent security incidents related to Service Account keys. Alternatively, your organization may have been automatically enforced with the policy as part of Secure by Default enforcements. 

Recommended Next Steps: Service account keys are a security risk if not managed correctly. You should choose a more secure alternative  whenever possible. If you must authenticate with a service account key, you or your administrator will need to disable  the iam.disableServiceAccountKeyCreation constraint and then retry creating the service account key.

Verfolgungsnummer: c5743236318768471

---
disable and enable service account keys:
https://cloud.google.com/iam/docs/keys-disable-enable

gcloud iam service-accounts keys enable KEY_ID \
    --iam-account=SA_NAME@PROJECT_ID.iam.gserviceaccount.com\
    --project=PROJECT_ID

key_id: 
sa_name: 
project_id: 
https://stackoverflow.com/questions/51508084/permission-denied-when-creating-a-service-account-key-using-the-gcloud-cli
---
https://cloud.google.com/sdk/docs/install

## Webhooks
### Trello
https://developer.atlassian.com/cloud/trello/guides/rest-api/webhooks/

https://developer.atlassian.com/cloud/trello/guides/rest-api/action-types/


#### Included Action Types
acceptEnterpriseJoinRequest
addAttachmentToCard
addChecklistToCard
addMemberToBoard
addMemberToCard
addMemberToOrganization
addOrganizationToEnterprise
addToEnterprisePluginWhitelist
addToOrganizationBoard
commentCard
convertToCardFromCheckItem
copyBoard
copyCard
copyCommentCard
createBoard
createCard
createList
createOrganization
deleteBoardInvitation
deleteCard
deleteOrganizationInvitation
disableEnterprisePluginWhitelist
disablePlugin
disablePowerUp
emailCard
enableEnterprisePluginWhitelist
enablePlugin
enablePowerUp
makeAdminOfBoard
makeNormalMemberOfBoard
makeNormalMemberOfOrganization
makeObserverOfBoard
memberJoinedTrello
moveCardFromBoard
moveCardToBoard
moveListFromBoard
moveListToBoard
removeChecklistFromCard
removeFromEnterprisePluginWhitelist
removeFromOrganizationBoard
removeMemberFromCard
removeOrganizationFromEnterprise
unconfirmedBoardInvitation
unconfirmedOrganizationInvitation
updateBoard
updateCard
updateCheckItemStateOnCard
updateChecklist
updateList
updateMember
updateOrganization

#### Excluded Action Types
addAdminToBoard (Deprecated in favor of makeAdminOfBoard)
addAdminToOrganization (Deprecated in favor of makeAdminOfOrganization)
addLabelToCard
copyChecklist
createBoardInvitation
createBoardPreference
createCheckItem
createLabel
createOrganizationInvitation
deleteAttachmentFromCard (Excluded from /lists/{listId}/actions)
deleteCheckItem
deleteComment
deleteLabel
makeAdminOfOrganization
removeAdminFromBoard (Deprecated in favor of makeNormalMemberOfBoard)
removeAdminFromOrganization (Deprecated in favor of makeNormalMemberOfOrganization)
removeLabelFromCard
removeMemberFromBoard
removeMemberFromOrganization
updateCheckItem
updateComment
updateLabel
voteOnCard
