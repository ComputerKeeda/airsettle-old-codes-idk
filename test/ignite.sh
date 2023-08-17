ignite scaffold type batch batch_number:uint merkle_root_hash prev_merkle_root_hash zk_proof
ignite scaffold type exelayer validator:array.string voting_power:array.uint latest_batch:uint latest_merkle_root_hash verification_key chain_info id creator

ignite scaffold message add_execution_layer verification_key chain_info --response id
ignite scaffold query show_execution_layer id --response exelayer:Exelayer
ignite scaffold query list_execution_layers --response exelayer:Exelayer --paginated

ignite scaffold message add_batch id batch_number:uint merkle_root_hash prev_merkle_root_hash zk_proof --response batch_status:bool
ignite scaffold query show_batch id batch_number:uint --response batch:Batch

ignite scaffold type vkey id verification_key
ignite scaffold type exelayer_chains creator id:array.string
ignite scaffold query chain_list creator_address --response exelayer_chains:array.string
ignite scaffold query chain_list_detailed creator_address --response chain:Exelayer --paginated
ignite scaffold query verification_key id --response vkey
ignite scaffold query verify id batch_number:uint inputs --response result:bool,message

ignite scaffold type poll poll_id chain_id new_validator votes_done_by:array.string votes:array.string total_validators:uint is_complete:bool start_date
ignite scaffold message add_validator new_validator_address chain_id --response voting_poll_id
ignite scaffold query list_add_validators_polls --response poll_ids:array.string --paginated
ignite scaffold query validator_poll_details poll_id --response poll:Poll
# ? Testing
ignite scaffold message submit_validator_vote poll_id vote:bool --response success:bool,poll_result,message,description
# TODO: 
# ignite scaffold query list_polls chainid --response poll:Poll --paginated

ignite scaffold type rollbackpoll
# only one rollback at a time. [approx all validators will target same batchnumber in case of wrong batch]
ignite scaffold message req_rollback batchnumber:uint chainid --response success:bool,message
ignite scaffold query get_rollback_status chainid --response rollbackpoll:Rollbackpoll # rollback in this batch
# until this is resolve no neighter new batch, nor new rollback request can be submitted
ignite sacffold message vote_rollback chainid batchnumber:uint vote:bool --response success:bool,message